package monerocnutils

import (
	"fmt"
	"testing"
)

const onlyMinerBlockTemplate = "0b0be2e4daec05e1c0a4a7bb5b3b658b518993017090bb110fe09402be9f07fded36391de176af0000000002a1f47401ffe5f3740192f0ffe7e545023306faf1cb524c328257f9d2ad35f78c4bb6ef0167ca60339aff0abb397a73803401c251fe844c41d6e3429efe40d017c7963bb15a4816c6ff611bf8cf3e0ac7bd5702110000000000000000000000000000000000000cd9ab785badfc78eda6ba4d7c2bd279ed450b94ab04e98a5c90ed7a989e8aae93b0c46214d0d68b23af4b7a3ab8ed6cfda825206003cbfa5d4810397a7da13f6c91b3c8617c561a11eea018bb5551cdb7c9552dbc475076e9be8740f5d462761f15d74b504dab729e9bb327e2f0d9c7ae201aa18efe8caee228269bf1c95ee6facc96ef0e2c233165a0193dbbe0fc256ac3511edbe3a981bd5d0541aa2e2b87887e1ccd4ef4e3d4355e25bdbe9ed21e4c2ab599b4d117612caf8c979ba76436413d290aa5a71339a5ccb9c2ae53798dd2d198b7415847277ea398da34fb913b7e6a42a9103d82c4ef2817d1ec47d746b6a3bb81fdaf83d35b0c49e1d6bc6aa69c6b347b1ebaa3ca077a5847a1500f17dba525f41323e3d696b36a2741146ec65e8fae1a4b38bf7dfabe511b964129117bd6d582235c0c2829e0f2fadf51979d9c6ff750bf38250cfd882d3ef7accc9e80b154187a541b0e2be6f9b6096a632c95bd50cec3019496c01345491a1f9aef4ac86b2b1148339667a05fa60e2a0d978f"
const onlyMinerBlockTemplateHashingBlob = "0b0be2e4daec05e1c0a4a7bb5b3b658b518993017090bb110fe09402be9f07fded36391de176af0000000035b8d4dd8cee2f82dcece2dbea5afbdb4d41725c4f5a2d5a0c343921ab7c1e160d"

/*
728-{.
729-  "id": "0",.
730-  "jsonrpc": "2.0",.
731-  "result": {.
732-    "blockhashing_blob": "0b0be2e4daec05e1c0a4a7bb5b3b658b518993017090bb110fe09402be9f07fded36391de176af0000000035b8d4dd8cee2f82dcece2dbea5afbdb4d41725c4f5a2d5a0c343921ab7c1e160d",.
733:    "blocktemplate_blob": "0b0be2e4daec05e1c0a4a7bb5b3b658b518993017090bb110fe09402be9f07fded36391de176af0000000002a1f47401ffe5f3740192f0ffe7e545023306faf1cb524c328257f9d2ad35f78c4bb6ef0167ca60339aff0abb397a73803401c251fe844c41d6e3429efe40d017c7963bb15a4816c6ff611bf8cf3e0ac7bd5702110000000000000000000000000000000000000cd9ab785badfc78eda6ba4d7c2bd279ed450b94ab04e98a5c90ed7a989e8aae93b0c46214d0d68b23af4b7a3ab8ed6cfda825206003cbfa5d4810397a7da13f6c91b3c8617c561a11eea018bb5551cdb7c9552dbc475076e9be8740f5d462761f15d74b504dab729e9bb327e2f0d9c7ae201aa18efe8caee228269bf1c95ee6facc96ef0e2c233165a0193dbbe0fc256ac3511edbe3a981bd5d0541aa2e2b87887e1ccd4ef4e3d4355e25bdbe9ed21e4c2ab599b4d117612caf8c979ba76436413d290aa5a71339a5ccb9c2ae53798dd2d198b7415847277ea398da34fb913b7e6a42a9103d82c4ef2817d1ec47d746b6a3bb81fdaf83d35b0c49e1d6bc6aa69c6b347b1ebaa3ca077a5847a1500f17dba525f41323e3d696b36a2741146ec65e8fae1a4b38bf7dfabe511b964129117bd6d582235c0c2829e0f2fadf51979d9c6ff750bf38250cfd882d3ef7accc9e80b154187a541b0e2be6f9b6096a632c95bd50cec3019496c01345491a1f9aef4ac86b2b1148339667a05fa60e2a0d978f",.
734-    "difficulty": 41893124737,.
735-    "expected_reward": 2398152030226,.
736-    "height": 1915365,.
737-    "prev_hash": "e1c0a4a7bb5b3b658b518993017090bb110fe09402be9f07fded36391de176af",.
738-    "reserved_offset": 128,.
739-    "status": "OK",.
740-    "untrusted": false.
741-  }.
*/

const minerTXBlockTemplate2 = "0b0bd999dbec05cd454db5b6ce8e2a7e75de1da6de3d27f532f2f7f91d12b5773b4c99982926b40000000002ea9d7601ffae9d7601cbd8e4f6fe42025b080825f9927a1f30bee9dfdb65361c6e88ed35790f11fe4c2501aa02e206ff3401c13e76ce528d2dab6f88dde278170b514ea47b752fa7a09fd1ba9ee175a92ebe02110000000000000000000000000000000000001167747f5e29cb6f4c225724082c2d77bb7d6d11dcdeac51f26367fb16d2e131b2af45a0b0ff5c65e630c3c8f38b431824a9b6a0188132d9105d71e2a249e45b4c569eeff2393cf94ffddef434449e1078d51cbe1ae4511ef5dd343e59da548e95504310e7807bf9ea71c2a2e6ad5c8f326b6292bf305a9cc97da9a79d497f378282d87b1c4992eba1ec5d7b76eb373b15669cc38c077615f1a5466c551c28b2a5baf9e6fdf7e5af8b2d07c716e8bfa2a5e6611bc753dc52d312c0a586ea24278c2a2adc4aed90daf324ccab11d4adc562575b331a02f9199bbd0cd10c68745bbb245fd37b23d31c7a73e739261982a15fe525d3f969508a524de00de15f13239756a68d8f43545f7d0a7b5c9a8a89ad8e8b0192a00b7a547987808e830468398439d8e0e6194f1d451f930baa776ae9719ab6fe12b4574f1c70565bb4ae755562089c0307ffa74c556024ca26873d7746070f118f485fd2394a38bda0e9dee9d0455ab7c0f3834a777657305fbd4985e0740a6b0e9d0da0f6a0ba8a7b4fda75a75109fc3d372fffb8a3282427c6748c2f4472fe4e13d22b915e1b19688c913ea9a8dcbcaeb6d96df0b1d8f18de9a5cfef8bb12c981aaba3c5ad8767f326308c299cbfa28521b2a6edeaf1e83b9c8824c74db31bb89678ed18c2b9190e0766d3f84b639972aff8fb188f51094c08c6a35057db364348bec656aeab4aa63fc27caf3aa5b4cc99d23f8c126f543b2839d240a0a77ace00b3ebf4d834c604230ca530"
const minerTXBlockTemplate2HashingBlob = "0b0bd999dbec05cd454db5b6ce8e2a7e75de1da6de3d27f532f2f7f91d12b5773b4c99982926b40000000018a060e0b7d7d3db1e36c464e5a92be8520c4baecc69363e29cfdd40dd7140f612"

/*
147167-{.
147168-  "id": "0",.
147169-  "jsonrpc": "2.0",.
147170-  "result": {.
147171-    "blockhashing_blob": "0b0bd999dbec05cd454db5b6ce8e2a7e75de1da6de3d27f532f2f7f91d12b5773b4c99982926b40000000018a060e0b7d7d3db1e36c464e5a92be8520c4baecc69363e29cfdd40dd7140f612",.
147172:    "blocktemplate_blob": "0b0bd999dbec05cd454db5b6ce8e2a7e75de1da6de3d27f532f2f7f91d12b5773b4c99982926b40000000002ea9d7601ffae9d7601cbd8e4f6fe42025b080825f9927a1f30bee9dfdb65361c6e88ed35790f11fe4c2501aa02e206ff3401c13e76ce528d2dab6f88dde278170b514ea47b752fa7a09fd1ba9ee175a92ebe02110000000000000000000000000000000000001167747f5e29cb6f4c225724082c2d77bb7d6d11dcdeac51f26367fb16d2e131b2af45a0b0ff5c65e630c3c8f38b431824a9b6a0188132d9105d71e2a249e45b4c569eeff2393cf94ffddef434449e1078d51cbe1ae4511ef5dd343e59da548e95504310e7807bf9ea71c2a2e6ad5c8f326b6292bf305a9cc97da9a79d497f378282d87b1c4992eba1ec5d7b76eb373b15669cc38c077615f1a5466c551c28b2a5baf9e6fdf7e5af8b2d07c716e8bfa2a5e6611bc753dc52d312c0a586ea24278c2a2adc4aed90daf324ccab11d4adc562575b331a02f9199bbd0cd10c68745bbb245fd37b23d31c7a73e739261982a15fe525d3f969508a524de00de15f13239756a68d8f43545f7d0a7b5c9a8a89ad8e8b0192a00b7a547987808e830468398439d8e0e6194f1d451f930baa776ae9719ab6fe12b4574f1c70565bb4ae755562089c0307ffa74c556024ca26873d7746070f118f485fd2394a38bda0e9dee9d0455ab7c0f3834a777657305fbd4985e0740a6b0e9d0da0f6a0ba8a7b4fda75a75109fc3d372fffb8a3282427c6748c2f4472fe4e13d22b915e1b19688c913ea9a8dcbcaeb6d96df0b1d8f18de9a5cfef8bb12c981aaba3c5ad8767f326308c299cbfa28521b2a6edeaf1e83b9c8824c74db31bb89678ed18c2b9190e0766d3f84b639972aff8fb188f51094c08c6a35057db364348bec656aeab4aa63fc27caf3aa5b4cc99d23f8c126f543b2839d240a0a77ace00b3ebf4d834c604230ca530",.
147173-    "difficulty": 39711778240,.
147174-    "expected_reward": 2301814713419,.
147175-    "height": 1937070,.
147176-    "prev_hash": "cd454db5b6ce8e2a7e75de1da6de3d27f532f2f7f91d12b5773b4c99982926b4",.
147177-    "reserved_offset": 128,.
147178-    "status": "OK",.
147179-    "untrusted": false.
147180-  }.
147181-}

*/

func TestParseBlockFromTemplateBlob(t *testing.T) {
	_, _ = ParseBlockFromTemplateBlob(onlyMinerBlockTemplate)
	_, _ = ParseBlockFromTemplateBlob(minerTXBlockTemplate2)
}

func TestGetBlockHashingBlob(t *testing.T) {
	b, _ := ParseBlockFromTemplateBlob(onlyMinerBlockTemplate)
	blob, _ := GetBlockHashingBlob(b)
	fmt.Printf("%x\n", blob)
	if fmt.Sprintf("%x", blob) != onlyMinerBlockTemplateHashingBlob {
		t.Fatal("Failed to properly convert blob back to hashing blob in comparison with the stored one")
	}
}
