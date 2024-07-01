/*
 *        _____   _          __  _____   _____   _       _____   _____
 *      /  _  \ | |        / / /  _  \ |  _  \ | |     /  _  \ /  ___|
 *      | | | | | |  __   / /  | | | | | |_| | | |     | | | | | |
 *      | | | | | | /  | / /   | | | | |  _  { | |     | | | | | |   _
 *      | |_| | | |/   |/ /    | |_| | | |_| | | |___  | |_| | | |_| |
 *      \_____/ |___/|___/     \_____/ |_____/ |_____| \_____/ \_____/
 *
 *  Copyright (c) 2023 by OwOTeam-DGMT (OwOBlog).
 * @Date         : 2024-02-04 14:10:45
 * @Author       : HanskiJay
 * @LastEditors  : HanskiJay
 * @LastEditTime : 2024-06-10 01:31:56
 * @E-Mail       : support@owoblog.com
 * @Telegram     : https://t.me/HanskiJay
 * @GitHub       : https://github.com/Tommy131
 */
package taskify

import (
	"log"
	"os"
	"owoweb/utils"
)

const (
	updateLogPath      = utils.STORAGE_PATH + "logs/update.log"
	bugLogPath         = utils.STORAGE_PATH + "logs/bug.log"
	defaultUpdateMsg   = "Optimized UI overflow bug & New Logo designed"
	defaultDownloadURL = "https://github.com/Tommy131/Taskify/releases"
)

var (
	updateLog *log.Logger
	bugLog    *log.Logger
)

func init() {
	updateLogFile, _ := os.OpenFile(updateLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	bugLogFile, _ := os.OpenFile(bugLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	updateLog = log.New(updateLogFile, "UPDATE: ", log.Ldate|log.Ltime|log.Lshortfile)
	bugLog = log.New(bugLogFile, "BUG REPORT: ", log.Ldate|log.Ltime|log.Lshortfile)
}
