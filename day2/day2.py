A = 1
B = 2
C = 3

X = 1
Y = 2
Z = 3

with open("input.txt") as f:
  file = f.readlines()

  matches = []
  
  for match in file:
    match = match.split(" ")
    opponent = match[0]
    if opponent == 'A':
      opponent = 1
    elif opponent == 'B':
      opponent = 2
    else:
      opponent = 3

    me = match[1].strip("\n")
    if me == 'X':
      me = 'lose'
    elif me == 'Y':
      me = 'draw'
    elif me == 'Z':
      me = 'win'
    
    matches.append((opponent, me))

    
  #print(matches)

  score = 0
  for (op,me) in matches:
    if me == 'lose':
      if op == 1:
        score += 3
      else:
        score += op-1
    elif me == 'draw':
      score += op + 3
    elif me == 'win':
      if op == 3:
        score += 7
      else:
        score += op + 1 + 6
        
  print(score)
    