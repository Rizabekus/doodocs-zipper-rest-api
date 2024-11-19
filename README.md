# Zipper/Archive Service REST API

## Overview

Zipper/Archive Service REST API is a tool for managing file archives. It provides interface for workign with archive files, converting files into arhive as well as an opportunity to send pdf/doc fiels to emails.

## Configs

Before Usage please fill out configs in .env file with your data.

## Usage

* To run this app:
``` 
go run ./cmd/web

```

## REST Methods

1. **Get Archive Information**
   - Endpoint: `/api/archive/information`
   - Method: `POST`
   - Retrieves data given an archive file through "multipart/form-data" in the section "files".

2. **Convert Files Into Archive**
   - Endpoint: `/api/archive/files`
   - Method: `POST`
   - Converts files of format xml,docx,jpeg,png into archive

3. **Send email**
   - Endpoint: `/api/archive/mail`
   - Method: `POST`
   - Sends to given emails in "emails" sections and a pdf/docx file in "file" section


## Logging

The code is extensively covered with debug and info logs to facilitate troubleshooting and monitoring.

## Configuration

Sensitive configuration data is stored in a `.env` file.

