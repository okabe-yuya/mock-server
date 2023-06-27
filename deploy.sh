gcloud config list | grep "project"

read -p "Is it okay to deploy to Cloud Run with these settings?(y/n)" yn
  case $yn in
      [Yy]* ) gcloud run deploy;;
      [Nn]* ) exit;;
      * ) echo "Please answer yes or no.";;
  esac
