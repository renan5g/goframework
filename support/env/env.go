package env

import "runtime"

// IsWindows returns whether the current operating system is Windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux returns whether the current operating system is Linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsDarwin returns whether the current operating system is Darwin.
func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

// IsArm returns whether the current CPU architecture is ARM.
func IsArm() bool {
	return runtime.GOARCH == "arm" || runtime.GOARCH == "arm64"
}

// IsX86 returns whether the current CPU architecture is X86.
func IsX86() bool {
	return runtime.GOARCH == "386" || runtime.GOARCH == "amd64"
}

// Is64Bit returns whether the current CPU architecture is 64-bit.
func Is64Bit() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
}
