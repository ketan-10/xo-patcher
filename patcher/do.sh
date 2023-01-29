~/go/bin/go-fanout --command="/home/ketan/go/bin/goimports -w" --chunk=5 -- xo_gen/enum/*
~/go/bin/go-fanout --command="/home/ketan/go/bin/goimports -w" --chunk=5 -- xo_gen/table/*
~/go/bin/go-fanout --command="/home/ketan/go/bin/goimports -w" --chunk=5 -- xo_gen/repo/*
~/go/bin/go-fanout --command="/home/ketan/go/bin/goimports -w" --chunk=5 -- xo_gen/xo_wire/*


# go install github.com/google/wire/cmd/wire@65ae46b7eaa146e99673e290251ea26f28139362
# if you install Above version of wire, it's old version which does not need to check pointer.
# /home/ketan/go/bin/wire ./wire_app
