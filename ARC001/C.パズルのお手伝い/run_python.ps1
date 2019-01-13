param($py_source, $testDataFile)

$CommandName = Split-Path -Leaf $PSCommandPath

if (-Not($py_source) -Or -Not($testDataFile)) {
    Write-Host "Usage : $CommandName py_source testDataFile"
    exit
}

Get-Content $testDataFile | python $py_source
