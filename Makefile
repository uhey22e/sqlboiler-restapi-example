OPENAPI_YML := openapi.yml

OAPI_PACKAGE := restapi
SQLBOILER_PACKAGE := boiler

.PHONY: generate
generate: $(SQLBOILER_PACKAGE) $(OAPI_PACKAGE)/types.go $(OAPI_PACKAGE)/server.go
$(SQLBOILER_PACKAGE):
	sqlboiler -o $(SQLBOILER_PACKAGE) -p $(SQLBOILER_PACKAGE) psql
$(OAPI_PACKAGE):
	mkdir -p $@
$(OAPI_PACKAGE)/types.go: $(OPENAPI_YML) restapi
	oapi-codegen -generate types -package $(OAPI_PACKAGE) -o $@.tmp $<
	gomodifytags -all -add-tags boil -transform snakecase -all -file $@.tmp > $@
	rm -f $@.tmp
$(OAPI_PACKAGE)/server.go: $(OPENAPI_YML) restapi
	oapi-codegen -generate chi-server -package $(OAPI_PACKAGE) -o $@ $<

clean:
	rm -rf $(SQLBOILER_PACKAGE)
	rm -rf $(OAPI_PACKAGE)
