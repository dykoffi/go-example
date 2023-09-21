export SONARQUBE_URL=https://system.sonarqube.eneci.net
export YOUR_PROJECT_KEY=ciql
export YOUR_REPO=.

docker run \
    --rm \
    --network host \
    -e SONAR_HOST_URL="${SONARQUBE_URL}" \
    -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=${YOUR_PROJECT_KEY}" \
    -e SONAR_TOKEN="sqp_9123c756fa58c7a43da6694f293589a09cc9d0fa" \
    -v "${YOUR_REPO}:/usr/src" \
    sonarsource/sonar-scanner-cli 