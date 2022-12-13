from functools import cmp_to_key

def packets_cmp(l, r):
  if evalPackets(l,r):
    return -1
  return 1

def evalPackets(left, right) -> bool:
  if len(left) == 0 and len(right) > 0:
    return True
  elif len(right) == 0 and len(left) > 0:
    return False
  elif len(left) == 0 and len(right) == 0:
    return
  if type(left[0]) == int and type(right[0]) == int:
    if left[0] < right[0]:
      return True
    elif left[0] == right[0]:
      return evalPackets(left[1:], right[1:])
    else:
      print("how did we end up here")
      return False
  
  if type(left[0]) == list and type(right[0]) == list:
    #print("list case", left[0], right[0])
    res = evalPackets(left[0], right[0])
    if res == None:
      return evalPackets(left[1:], right[1:])
    else:
      return res
  
  if type(left[0]) == int and type(right[0]) == list:
    left[0] = [left[0]]
    return evalPackets(left, right)
  elif type(right[0]) == int and type(left[0]) == list:
    right[0] = [right[0]]
    return evalPackets(left, right)
  
  evalPackets(left[1:], right[1:])
  #print("finding end")
  #return False    

i = 1
sum = 0
packets = []
with open("input.txt") as f:
  pairs = f.read().split("\n\n")
  for pair in pairs:
    p = pair.split("\n")
    left = p[0]
    right = p[1]    
    left = eval(left)
    right = eval(right)
    print(left)
    print(right)
    res = evalPackets(left, right)
    if res:
      sum += i
    i += 1
    packets.append(left)
    packets.append(right)
  
  diva = eval("[[2]]")
  packets.append(diva)
  divb = eval("[[6]]")
  packets.append(divb)
  cmp_key = cmp_to_key(packets_cmp)
  packets.sort(key=cmp_key)
  div1 = packets.index(diva)
  div2 = packets.index(divb)
  print(sum)
  print((div1+1) * (div2+1))