Write-Host "PreData"
python .\build_random_testdata.py | tee aaa.txt
Get-Content aaa.txt | python .\ARC001C.py | python .\result_check.py
