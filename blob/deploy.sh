az group create -n "imagesrg" --location "eastasia"
az deployment group create -n "imagesdeploy" -g "imagesrg" --template-file ./deploy.json --parameters ./deploy.parameters.json