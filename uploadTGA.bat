docker build . -f portrait.dockerfile -t nwnci/portrait
docker run --rm -it --env-file %cd%/env.list nwnci/portrait:latest