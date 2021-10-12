#!/bin/sh

PROJECT_ID=$1
if [ -z "$PROJECT_ID" ]; then
  echo "Please specify the ProjectID."
  exit 1
fi

SERVICE=$2
if [ -z "$SERVICE" ]; then
  SERVICE="default"
fi

echo "Project=$PROJECT_ID, Service=$SERVICE"

ALL_NO_TRAFFIC_VERSIONS=$(gcloud app versions list --format="value(id)" --service=$SERVICE --filter="traffic_split=0" --sort-by="~lastDeployedTime" --project=$PROJECT_ID)
echo "$ALL_NO_TRAFFIC_VERSIONS"

if test "$(echo "$ALL_NO_TRAFFIC_VERSIONS" | wc -l)" -gt 30; then
  DELETING_VERSIONS=$(echo "$ALL_NO_TRAFFIC_VERSIONS" | tail -n +21)
  gcloud app versions delete --quiet --project=$PROJECT_ID --service=$SERVICE $(echo "$DELETING_VERSIONS")
fi
