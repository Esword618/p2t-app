/*
Copyright (c) GitHub, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// These paths are all drawn to a 10x10 view box and replicate the symbols
// seen on Windows 10 window controls.
export const closePath =
    ['M 0,0 0,0.7 4.3,5 0,9.3 0,10 0.7,10 5,5.7 9.3,10 10,10 10,9.3 5.7,5 10,0.7 10,0 9.3,0 5,4.3 0.7,0 Z']
export const restorePath =
    ['m 2,1e-5 0,2 -2,0 0,8 8,0 0,-2 2,0 0,-8 z m 1,1 6,0 0,6 -1,0 0,-5 -5,0 z m -2,2 6,0 0,6 -6,0 z']
export const maximizePath = ['M 0,0 0,10 10,10 10,0 Z M 1,1 9,1 9,9 1,9 Z']
export const minimizePath = ['M 0,5 10,5 10,6 0,6 Z']
export const maximizePath_win11 = [
    `M1.3,0c0.1,0,0.3,0,0.4,0C4,0,6.2,0,8.4,0c0.4,0,0.8,0.2,1.1,0.5C9.9,0.8,10,1.2,10,1.6c0,2.3,0,4.6,0,6.9
    c0,0.4-0.2,0.9-0.5,1.1C9.2,9.9,8.8,10,8.4,10c-2.1,0-4.3,0-6.4,0c-0.3,0-0.6,0-0.9-0.1c-0.4-0.1-0.8-0.4-1-0.8C0,8.9,0,8.7,0,8.4
    c0-2.3,0-4.6,0-6.8c0-0.4,0.1-0.8,0.4-1.1C0.7,0.2,1,0.1,1.3,0 M1.4,0.9c-0.3,0-0.6,0.3-0.6,0.7c0,2.3,0,4.6,0,6.9
    c0,0.4,0.3,0.5,0.4,0.6c0.2,0.1,0.2,0,0.5,0c0.2,0,4.5,0,6.7,0c0.1,0,0.2,0,0.4-0.1C8.9,9,9.2,8.8,9.2,8.5c0-2.3,0-4.6,0-6.9
    c0-0.2-0.1-0.5-0.3-0.6C8.8,0.9,8.6,0.9,8.4,0.9c-2.2,0-4.5,0-6.7,0C1.6,0.9,1.6,0.9,1.4,0.9z`,
]
export const restorePath_win11 = [
    `M2,1.5C2.1,0.7,2.8,0,3.6,0C5,0,6.4,0,7.8,0C9,0.1,10,1.1,10,2.3c0,1.4,0,2.8,0,4.1C10,7.2,9.3,7.9,8.5,8
    c0-0.3,0-0.5,0-0.8c0.4,0,0.8-0.4,0.7-0.9c0-1.3,0-2.6,0-3.9c0.1-0.9-0.7-1.7-1.5-1.7c-1.4,0-2.7,0-4.1,0c-0.4,0-0.8,0.3-0.9,0.7
    C2.5,1.5,2.3,1.5,2,1.5z`,
    `M1.5,2c1.7,0,3.3,0,5,0C7.2,2,8,2.7,8,3.6c0,1.6,0,3.3,0,4.9C8,9.3,7.2,10,6.4,10c-1.6,0-3.2,0-4.8,0C0.7,10,0,9.3,0,8.4
    c0-1.6,0-3.1,0-4.7C-0.1,2.9,0.6,2.1,1.5,2 M1.5,2.8C1,2.8,0.7,3.3,0.8,3.8c0,1.5,0,3.1,0,4.6c0,0.5,0.4,0.9,0.9,0.9
    c1.5,0,3.1,0,4.6,0c0.5,0.1,1-0.3,1-0.9c0-1.6,0-3.2,0-4.8c0-0.5-0.5-0.9-1-0.8C4.7,2.8,3.1,2.7,1.5,2.8z`
]

export const onTopPath = ['M927.701333 375.978667l-279.637333-279.637334A37.546667 37.546667 0 0 0 621.397333 85.333333a37.546667 37.546667 0 0 0-26.666666 11.008L411.904 279.296a366.506667 366.506667 0 0 0-41.770667-2.304 374.016 374.016 0 0 0-234.325333 82.048 37.717333 37.717333 0 0 0-3.072 56.064l206.208 206.208-244.48 244.224a17.92 17.92 0 0 0-5.205333 11.093333l-3.84 42.24a18.090667 18.090667 0 0 0 18.048 19.754667c0.554667 0 1.109333 0 1.706666-0.128l42.24-3.84a17.92 17.92 0 0 0 11.093334-5.248l244.437333-244.437333 206.208 206.208a37.546667 37.546667 0 0 0 26.666667 11.008 37.546667 37.546667 0 0 0 29.397333-14.08 374.826667 374.826667 0 0 0 79.658667-276.224l182.826666-182.826667a37.589333 37.589333 0 0 0 0-53.077333z m-240.725333 178.346666l-27.776 27.818667 4.266667 39.04a294.997333 294.997333 0 0 1-34.474667 174.677333L228.309333 394.922667a293.546667 293.546667 0 0 1 174.634667-34.517334l39.04 4.309334 27.818667-27.776 151.722666-151.722667 217.301334 217.301333-151.850667 151.850667z']


export const svgContainer = (sth: string, size?: { width: number, height: number }, style?: string) => `<svg style="${style}" width="${size?.width || 10}" height="${size?.height || 10}"><path d="${sth}" /></svg>`;