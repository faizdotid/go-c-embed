package main

// #include <Windows.h>
import "C"
import (
	"go-c-embed/utils"
	"unsafe"
)

// CreateMessageBox creates a message box with the given message and title.
func CreateMessageBox(s, title string) int {
	message := C.CString(s)
	defer C.free(unsafe.Pointer(message))

	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	return int(C.MessageBox(
		nil,
		message,
		cTitle,
		utils.MB_ICONWARNING,
	))
}

func main() {
	className := C.CString("BUTTON")
	defer C.free(unsafe.Pointer(className))

	buttonText := C.CString("Click Me")
	defer C.free(unsafe.Pointer(buttonText))

	hwnd := C.CreateWindowExA(
		C.DWORD(0),
		className,
		buttonText,
		C.WS_OVERLAPPEDWINDOW,
		C.CW_USEDEFAULT,
		C.CW_USEDEFAULT,
		C.CW_USEDEFAULT,
		C.CW_USEDEFAULT,
		nil,
		nil,
		nil,
		nil,
	)

	C.ShowWindow(hwnd, C.SW_SHOWDEFAULT)
	
	CreateMessageBox("Hello, World!", "Hello")
}
