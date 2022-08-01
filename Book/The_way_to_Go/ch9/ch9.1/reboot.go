package main

const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
const LINUX_REBOOT_MAGIC2 uintptr = 672274793
const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

//Syscall 已弃用
func main() {
	//SYS.REBOOT 已经弃用 go-version 1.18
	//syscall.Syscall(syscall.SYS_REBOOT,
	//	LINUX_REBOOT_MAGIC1,
	//	LINUX_REBOOT_MAGIC2,
	//	LINUX_REBOOT_CMD_RESTART)
}
