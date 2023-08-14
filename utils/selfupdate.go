package utils

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

//func init() {
//	// initialize update.log file and set log output to file
//	file, err := os.OpenFile("./update.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	if err == nil {
//		log.SetOutput(file)
//	} else {
//		log.Info("Failed to log to file, using default stderr")
//	}
//
//	// Only log the warning severity or above.
//	log.SetLevel(log.InfoLevel)
//}

type Update struct {
	dagFolderPath    *string
	oldAppBinaryPath *string
	newAppBinaryPath *string
	appName          *string
	PID              *string
}

func NewUpdate() {
	var update Update
	update.dagFolderPath = flag.String("init_dag_path", "", "Enter the directory path to dag folder")
	update.oldAppBinaryPath = flag.String("old_app_path", "", "Enter the old app binary path")
	update.newAppBinaryPath = flag.String("new_app_path", "", "Enter the new app binary path")
	update.appName = flag.String("app_name", "", "Enter the app name")
	update.PID = flag.String("pid", "", "Enter the app PID")
	log.Infoln("dagFolderPath===>   ", update.dagFolderPath)
	log.Infoln("oldAppBinaryPath===>   ", update.oldAppBinaryPath)
	log.Infoln("newAppBinaryPath===>   ", update.newAppBinaryPath)
	log.Infoln("appName===>   ", update.dagFolderPath)
	log.Infoln("PID===>   ", update.PID)
	flag.Parse()
	update.Run()
}

func (u *Update) Run() {
	var err error
	// Clean up old update artifacts
	err = u.CleanUp()
	if err != nil {
		log.Fatalf("Unable to clear previous local state: %v", err)
	}
	//err = u.BackupApp()
	//if err != nil {
	//	log.Fatalf("Unable to Backup App: %v", err)
	//}
	err = u.TerminateAppService()
	if err != nil {
		log.Errorf("Unable to terminate App: %v", err)
		err = u.RestoreBackup()
		if err != nil {
			log.Fatal("Unable to restore backup: %v", err)
		}
	}
	err = u.ReplaceAppBinary()
	if err != nil {
		log.Errorf("Unable to overwrite old installation: %v", err)
		err = u.RestoreBackup()
		if err != nil {
			log.Fatalf("Unable to restore backup: %v", err)
		}
	}
	err = u.LaunchAppBinary()
	if err != nil {
		log.Errorf("Unable to overwrite old installation: %v", err)
		err = u.RestoreBackup()
		if err != nil {
			log.Fatalf("Unable to restore backup: %v", err)
		}
	}
}

// CleanUp removes uneccesary artifacts from the update process
func (u *Update) CleanUp() error {
	if fileExists(*u.dagFolderPath + "/backup") {
		err := os.RemoveAll(*u.dagFolderPath + "/backup")
		if err != nil {
			return err
		}
	}
	return nil
}

// BackupApp backs up the old binary in case of a failed update.
func (u *Update) BackupApp() error {
	// Create backup folder in ~/.dag
	err := os.Mkdir(*u.dagFolderPath+"/backup", 0755)
	if err != nil {
		return fmt.Errorf("unable to create backup folder: %v", err)
	}

	// Copy the old App binary into ~/.dag/backup/
	err = fileCopy(*u.oldAppBinaryPath, *u.dagFolderPath+"/backup/backup_"+*u.appName)
	if err != nil {
		return fmt.Errorf("unable to backup old App installation: %v", err)
	}

	return nil
}

// RestoreBackup restores the backed up files
func (u *Update) RestoreBackup() error {
	log.Infoln("Restoring Backup...")
	// Copy the old App binary from ~/.dag/backup/ to the old path
	//_, fileExt := getUserOS()
	err := fileCopy(*u.dagFolderPath+"/backup/backup_"+*u.appName, *u.oldAppBinaryPath)
	if err != nil {
		return fmt.Errorf("unable to overwrite old App binary: %v", err)
	}
	//// Copy update binary from ~/.dag/backup/update -> ~/.dag/update
	//if fileExists(*u.dagFolderPath + "/backup/update" + fileExt) {
	//	err = copy(*u.dagFolderPath+"/backup/update"+fileExt, *u.dagFolderPath+"/update"+fileExt)
	//	if err != nil {
	//		return fmt.Errorf("unable to copy update binary to .dag folder: %v", err)
	//	}
	//}

	log.Infoln("Backup successfully restored.")
	return nil

}

// TerminateAppService will terminate the application
func (u *Update) TerminateAppService() error {
	log.Infoln("关闭app")
	pid, _ := strconv.Atoi(*u.PID)
	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Panicf("%v", err)
	}
	err = proc.Kill()
	log.Infoln("关闭app end")
	if err != nil {
		return err
	}
	return nil
}
func (u *Update) ReplaceAppBinary() error {
	log.Infoln("替换app")
	// Replace old App binary with the new one
	//_, fileExt := getUserOS()
	err := fileCopy(*u.newAppBinaryPath, *u.oldAppBinaryPath)
	if err != nil {
		for i := 5; i > 0; i-- {
			time.Sleep(time.Duration(i) * time.Second)
			err = fileCopy(*u.newAppBinaryPath, *u.oldAppBinaryPath)
			if err == nil {
				break
			} else if err != nil && i == 0 {
				return fmt.Errorf("unable to overwrite old App binary: %v", err)
			}
		}
	}
	//// Replace old update binary with the new one
	//if fileExists(*u.newAppBinaryPath) {
	//	err = fileCopy(*u.newAppBinaryPath, *u.dagFolderPath+"/update"+fileExt)
	//	if err != nil {
	//		return fmt.Errorf("unable to copy update binary to .dag folder: %v", err)
	//	}
	//}
	log.Infoln("替换app end")
	return nil
}

// LaunchAppBinary executes the new App
func (u *Update) LaunchAppBinary() error {
	log.Infoln("重新吊起")
	log.Infoln("oldAppBinaryPath===>   ", u.oldAppBinaryPath)
	cmd := exec.Command(*u.oldAppBinaryPath)

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("unable to execute run command for new App: %v", err)
	}
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !os.IsNotExist(err)
}
func fileCopy(src string, dst string) error {
	// Read all content of src to data
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// Write data to dst
	err = os.WriteFile(dst, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

// getUserOS returns the users OS as well as the file extension of executables for said OS
func getUserOS() (string, string) {
	var osBuild string
	var fileExt string

	switch goos := runtime.GOOS; goos {
	case "darwin":
		osBuild = "darwin"
		fileExt = ""
	case "linux":
		osBuild = "linux"
		fileExt = ""
	case "windows":
		osBuild = "windows"
		fileExt = ".exe"
	default:
		osBuild = "unsupported"
		fileExt = ""
	}

	return osBuild, fileExt
}
