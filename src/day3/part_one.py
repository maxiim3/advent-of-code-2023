sample_file = [ln.replace("\n", "") for ln in open("sample.txt", "r").readlines()]
print(sample_file)

for ln in sample_file:
    ln.find(\\)
