mkdir drug-label

cd drug-label || exit
  curl -O https://download.open.fda.gov/drug/label/drug-label-0001-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0002-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0003-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0004-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0005-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0006-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0007-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0008-of-0009.json.zip
  curl -O https://download.open.fda.gov/drug/label/drug-label-0009-of-0009.json.zip
  unzip '*.json.zip'
cd - || exit

cd nsde || exit
  curl -O https://download.open.fda.gov/other/nsde/other-nsde-0001-of-0002.json.zip
  curl -O https://download.open.fda.gov/other/nsde/other-nsde-0002-of-0002.json.zip
  unzip '*.json.zip'
cd - || exit

go run main.go