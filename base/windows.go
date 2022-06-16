package base

/*
#include<stdio.h>
#include<windows.h>
#include<conio.h>

struct Coord {
	int X; //光标的X坐标
	int Y; //光标的Y坐标
};

struct Coord getCursorCoordinate()
{
	struct Coord coord;
    CONSOLE_SCREEN_BUFFER_INFO consoleBuffer;
   	HANDLE stdoutHandle = GetStdHandle(STD_OUTPUT_HANDLE);
    GetConsoleScreenBufferInfo(stdoutHandle, &consoleBuffer);
    coord.X = consoleBuffer.dwCursorPosition.X;
	coord.Y = consoleBuffer.dwCursorPosition.Y;
	return coord;
}

void setCursorCoordinate(int x,int y)
{
	COORD coord = {x, y};
	HANDLE handle = GetStdHandle(STD_OUTPUT_HANDLE);
	SetConsoleCursorPosition(handle, coord);
}

void setConsoleStatus(int status, int fore, int back)
{
	HANDLE handle = GetStdHandle(STD_OUTPUT_HANDLE);
	SetConsoleTextAttribute(handle, (WORD) (status | fore | back));
}

void hideCursor()
{
	CONSOLE_CURSOR_INFO cursor;
	cursor.bVisible = FALSE;
	cursor.dwSize = sizeof(cursor);
	HANDLE handle = GetStdHandle(STD_OUTPUT_HANDLE);
	SetConsoleCursorInfo(handle, &cursor);
}
*/
import "C"
import (
	"fmt"
)

//隐藏光标
func hideWindowsCursor() {
	C.hideCursor()
}

//获取光标坐标
func getWindowsCursorCoordinate() (int, int) {
	c := C.getCursorCoordinate()
	return int(c.X), int(c.Y)
}

//设置光标坐标
func setWindowsCursorCoordinate(x, y int) {
	C.setCursorCoordinate(C.int(x), C.int(y))
}

//设置状态
func setWindowsConsoleStatus(style, fore, back int) {
	C.setConsoleStatus(C.int(style), C.int(fore), C.int(back))
}

//**********************************************************************************************

// 打印到坐标
func printToWindowsToCoordinate(x, y int, msg string) {
	C.setCursorCoordinate(C.int(x), C.int(y))
	fmt.Print(msg)
}

// 打印带状态的字符串
func printToWindowsWithStatus(style, fore, back int, msg string) {
	C.setConsoleStatus(C.int(style), C.int(fore), C.int(back))
	fmt.Print(msg)
}

// 打印带状态的字符串到指定坐标
func printToWindowsWithStatusToCoordinate(x, y, style, fore, back int, msg string) {
	C.setCursorCoordinate(C.int(x), C.int(y))
	C.setConsoleStatus(C.int(style), C.int(fore), C.int(back))
	fmt.Print(msg)
}

// winForeBlack           = 0x0000
// winForeBlue            = 0x0001
// winForeGreen           = 0x0002
// WinForeLakeBlue        = 0x0003
// WinForeRed             = 0x0004
// WinForePurple          = 0x0005
// WinForeYellow          = 0x0006
// WinForeWhite           = 0x0007
// WinForeGrey            = 0x0008
// WinForeLightBlue       = 0x0009
// WinForeLightGreen      = 0x000A
// WinForeLightLightGreen = 0x000B
// WinForeLightRed        = 0x000C
// WinForeLightPurple     = 0x000D
// WinForeLightYellow     = 0x000E
// WinForeBrightWhite     = 0x000F

// winBackBlack           = 0x0000
// winBackBlue            = 0x0001
// winBackGreen           = 0x0002
// WinBackLakeBlue        = 0x0003
// WinBackRed             = 0x0004
// WinBackPurple          = 0x0005
// WinBackYellow          = 0x0006
// WinBackeWhite          = 0x0007
// WinBackGrey            = 0x0008
// WinBackLightBlue       = 0x0009
// WinBackLightGreen      = 0x000A
// WinBackLightLightGreen = 0x000B
// WinBackLightRed        = 0x000C
// WinBackLightPurple     = 0x000D
// WinBackLightYellow     = 0x000E
// WinBackBrightWhite     = 0x000F
