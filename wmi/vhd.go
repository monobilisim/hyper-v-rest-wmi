package wmi

func Vhd(vmName string) ([]byte, error) {
	ps := `Get-VHD -Id ` + vmName + ` | ConvertTo-Json`
	output, err := execPS(ps)
	return output, err
}
