# PDF Converter

Small HTTP service that uses Libre Office internally to convert most common file formats to PDF

## How to use

Send a POST request to the server containing the file that will be converted to PDF

```bash
curl -X POST \
-F "file=@myfile.docx" \
pdf-converter-address:3000 > /tmp/test.pdf
```

This service also includes a health cheking URL on root route

```bash
curl -v pdf-converter-address:3000
```