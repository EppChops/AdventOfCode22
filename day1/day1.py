

with open("input.txt") as f:
  file = f.readlines()

  cals = []
  count = 0
  for cal in file:
    if cal != '\n':
      count += int(cal)
      #print(count)
    else:
      cals.append(count)
      count = 0
  

  max1 = max(cals)
  cals.remove(max1)
  max2 = max(cals)
  cals.remove(max2)
  max3 = max(cals)

  total_max = max1 + max2 + max3
  print(total_max)