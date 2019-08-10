# nwn-portrait-digitalOcean-uploader

This is an application that I use in my CI for my personal nwn projects.

This will ingest all the .tga files you put into the tga folder, convert them to png. Then upload them to [Digital Ocean Spaces](https://cloud.digitalocean.com/spaces/).

Currently I have only tested this with DigitalOcean, but in concept this should work with AWS as well, no promises on any other providers either.

### example png:
https://nwn.urothis.dev/portrait/po_a_bird_rav_h.png

![Raven](https://nwn.urothis.dev/portrait/po_a_bird_rav_h.png)
![Raven](https://nwn.urothis.dev/portrait/po_a_bird_rav_l.png)
![Raven](https://nwn.urothis.dev/portrait/po_a_bird_rav_m.png)
![Raven](https://nwn.urothis.dev/portrait/po_a_bird_rav_s.png)
![Raven](https://nwn.urothis.dev/portrait/po_a_bird_rav_t.png)

# HOWTOUSE
1. Latest docker
2. To generate the access keys you need, follow [THIS](https://www.digitalocean.com/community/tutorials/how-to-create-a-digitalocean-space-and-api-key) guide
3. These values are all required, they are assigned in [env.list](https://github.com/urothis/nwn-portrait-digitalOcean-uploader/env.list)

Key | Value
------------ | -------------
endpoint | ENDPOINT
accessKeyID | ACCESS_KEY
secretAccessKey | SECRET_KEY
bucketName | SPACE_NAME default:"nwn"
subFolder | SUBFOLDER default:"portrait"
4. Run the relative .bat (windows) .sh (linux) for your os. 
5. Sit back and watch.
