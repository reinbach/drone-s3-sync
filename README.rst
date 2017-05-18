Drone S3 Sync Plugin
====================

`Drone <https://github.com/drone/drone>`_ plugin for syncing files to S3.

Build
-----

.. code-block:: bash

   GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o s3-sync
   docker build -t reinbach/drone-s3-sync .
   docker push reinbach/drone-s3-sync

Test
----

.. code-block:: bash

   docker run --rm \
     -e PLUGIN_ACCESS_KEY=<aws-access-key> \
     -e PLUGIN_SECRET_KEY=<aws-secret-key> \
     -e PLUGIN_SOURCE=path/from/ \
     -e PLUGIN_BUCKET=foo-bucket \
     -e PLUGIN_TARGET=/path/to/ \
     -e PLUGIN_REGION=us-east-1 \
     -e PLUGIN_ACL=public-read \
     -e PLUGIN_DELETE=true \
     reinbach/drone-s3-sync

Usage
-----

.. code-block:: yaml

    pipeline:
      s3-sync:
        image: reinach/drone-s3-sync
        source: path/from/
        bucket: "s3-bucket"
        target: /path/to/
        region: "us-east-1"
        acl: public-read
        delete: true
        secrets: [access_key, secret_key]
