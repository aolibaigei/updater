import json
import os


feed_dir = "./feed/"
out_dir = "./output/"
actve_cve_path = "/home/dev/gosrc/ubuntu-cve-tracker/active"
retired_cve_path = "/home/dev/gosrc/ubuntu-cve-tracker/retired"
nvd_files = []
active_cve = []
retired_cve = []

def add_file(outputpath,cveno,jsontext):
    path = outputpath + cveno + ".json"
    print(path)
    with open(path,'w') as f:
        f.write(str(jsontext).replace("'","\""))
        
    print("saved file "+ cveno)


def get_nvd_files(path):

    _files = []
    for parent, dirnames, filenames in os.walk(path):
        for filename in filenames:
            if ".json" in filename[-5:]:
                _files.append(filename)
    return _files



def get_active_cve(path):

    _files = []
    for parent, dirnames, filenames in os.walk(path):
        for filename in filenames:
            if "CVE-201" in filename:
                _files.append(filename)
    return _files


active_cve = get_active_cve(actve_cve_path)

with open("./feed/nvdcve-1.0-2019.json", 'r') as f:
        temp = json.loads(f.read())
        for i in temp["CVE_Items"]:
            cveno = i['cve']['CVE_data_meta']['ID']
            vendor_data = i['cve']['affects']['vendor']['vendor_data']
            # if len(vendor_data) > 0 :
            #     for x in range(0,len(vendor_data)):
            #         print(vendor_data[x])

            if len(vendor_data) > 0 and cveno in active_cve :
                add_file(out_dir,cveno,i)




