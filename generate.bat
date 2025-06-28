@echo off
oapi-codegen -config oapi-codegen.yaml api/openapi.yaml
echo.
echo Code generation complete.
pause