import json
import os

rootdir = "E:\\github\\ubuntu-cve-tracker\\active"
_files = []

list = os.listdir(rootdir)

for parent, dirnames, filenames in os.walk(rootdir):

    for filename in filenames:
        if "CVE-201" in filename:

            _files.append(filename)

        # with open("./feeds/nvdcve-1.0-2019.json", 'r') as f:
        #     temp = json.loads(f.read())

        #     for i in temp["CVE_Items"]:
        #         cveno = i['cve']['CVE_data_meta']['ID']
        #         print(i['cve']['affects']['vendor'])


print(_files)
