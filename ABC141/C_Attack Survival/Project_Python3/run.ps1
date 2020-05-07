param($file)

$CommandName = Split-Path -Leaf $PSCommandPath

if (-Not $file) {
    Write-Host "Usage: $CommandName <datafile>"
    exit
}

if (-Not(Test-Path $file)) {
    Write-Host "$file not found."
    exit
}


Get-Content $file

Write-Host "Get-Content $file | python .\ABC141_C_Attack_Survial.py"
Get-Content $file | python .\ABC141_C_Attack_Survial.py
