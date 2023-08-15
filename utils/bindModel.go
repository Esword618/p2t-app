package utils

import (
	"github.com/Esword618/p2t/global"
	"github.com/Esword618/p2t/model"
)

func (u *Utils) GetLatestDate() []model.LatexModel {

	// Get the latest updated time
	var latestModel model.LatexModel
	global.GlobalDb.Order("updated_at desc").First(&latestModel)

	latestDateStr := latestModel.UpdatedAt.Format("2006-01-02") // Format to YYYY-MM-DD

	// Fetch all records updated at the latest day
	var latestRecords []model.LatexModel
	global.GlobalDb.Where("date(updated_at) = ?", latestDateStr).Find(&latestRecords)

	return latestRecords

}

func (u *Utils) GetDataByTimeOrLatex(timeInput, latexInput string) []model.LatexModel {
	var latexModels []model.LatexModel
	query := global.GlobalDb

	if timeInput != "" {
		query = query.Where("date(updated_at) = ?", timeInput)
	}

	if latexInput != "" {
		query = query.Where("latex LIKE ?", "%"+latexInput+"%")
	}

	if err := query.Find(&latexModels).Error; err != nil {
		return nil
	}

	return latexModels
}

func (u *Utils) DeleteByUuid(uuid string) error {
	var latexModel model.LatexModel
	query := global.GlobalDb

	// Add the condition
	if uuid != "" {
		query = query.Where("uuid = ?", uuid)
	}

	// Delete the record
	if err := query.Delete(&latexModel).Error; err != nil {
		// Assuming you want to return the error to the caller
		return err
	}

	return nil
}
