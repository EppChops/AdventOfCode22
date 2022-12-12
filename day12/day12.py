import networkx as nx

G = nx.DiGraph()
with open("example.txt") as f:
  file = f.readlines()

  for i, line in enumerate(file):
    for j, c in enumerate(line):
      if c == "S":
        start = (i,j)
        file[i] = file[i].replace("S", "a")
      if c == "E":
        end = (i,j)
        file[i] = file[i].replace("E", "z")

  print(file[0])
  for i, line in enumerate(file):
    print(line)
    line = line.strip('\n')
    for j, c in enumerate(line):
      G.add_node((i, j))
      print(i)
      if c == "S":
        start = (i,j)
        c = "a"
      if c == "E":
        end = (i,j)
        c = "z"
      
      if j < len(line) - 1 and abs(ord(c) - ord(line[j+1])) <= 1:
        print(c, (i,j), (i, j+1), "fst")
        G.add_edge((i,j), (i, j+1))
        G.add_edge((i, j+1), (i,j))
      elif j < len(line) - 1 and ord(c) - ord(line[j+1]) > 1:
        print(c, (i,j))
        G.add_edge((i,j), (i, j+1))
      elif j < len(line) - 1 and ord(c) - ord(line[j+1]) <  -1:
        G.add_edge((i, j+1), (i,j))
      if j > 0 and abs(ord(c) - ord(line[j-1])) <= 1:
        G.add_edge((i,j), (i, j-1))
        G.add_edge((i, j-1), (i,j))
      elif  j > 0 and ord(c) - ord(line[j-1]) > 1:
        G.add_edge((i,j), (i, j-1))
      elif j > 0 and ord(c) - ord(line[j-1]) < -1:
        print(c, (i,j), "j-1")
        G.add_edge((i, j-1), (i,j))
      if i >0 and abs(ord(c) - ord(file[i-1][j])) <= 1:
        print(c, (i,j), (i-1, j), "Halp")
        G.add_edge((i,j), (i-1, j))
        G.add_edge((i-1,j), (i, j))
      elif i > 0 and ord(c) - ord(file[i-1][j]) > 1:
        print(c, (i,j), (i-1, j), "Halp5")
        G.add_edge((i,j), (i-1, j))
      elif i > 0 and ord(c) - ord(file[i-1][j]) < -1:
        G.add_edge((i-1,j), (i, j))
      if i < len(file) - 1 and abs(ord(c) - ord(file[i+1][j])) <= 1:
        print(c, (i,j), (i+1, j), "Halp2")
        G.add_edge((i,j), (i+1,j))
        G.add_edge((i+1,j), (i,j))
      elif i < len(file) - 1 and ord(c) - ord(file[i+1][j]) > 1:
        G.add_edge((i,j), (i+1,j))
        print(c, (i,j), (i+1, j), "Halp3")
      elif i < len(file) - 1 and ord(c) - ord(file[i+1][j]) < -1:
        G.add_edge((i+1,j), (i,j))
        print(c, (i,j), (i+1, j), "Halp4")
         
        """
        if ord(c)+1 <= ord(line[j+1]):
          print(ord(c), ord(line[j+1]), (i,j))
          G.add_edge((i, j), (i,j+1))
          if ord(line[j+1]) - ord(c) == 1 or  (ord(line[j+1]) - ord(c)) == 0:
            G.add_edge((i, j), (i, j+1))
      if i < len(file)-1 :
        if ord(c) + 1 <= ord(file[i+1][j]):
          G.add_edge((i+1, j), (i,j))
          if ord(file[i+1][j]) - ord(c) == 1 or  (ord(file[i+1][j]) - ord(c)) == 0:
            G.add_edge((i, j), (i+1, j))
      if j > 0:
        if ord(c)+1 <= ord(line[j-1]):
          G.add_edge((i, j-1), (i,j))
          if ord(line[j-1]) - ord(c) == 1 or  (ord(line[j-1]) - ord(c)) == 0:
            G.add_edge((i, j), (i, j-1))
      if i > 0:
        if ord(c) + 1 <= ord(file[i-1][j]):
          G.add_edge((i-1, j), (i,j))
          if ord(file[i-1][j]) - ord(c) == 1 or  (ord(file[i-1][j]) - ord(c)) == 0:
            G.add_edge((i, j), (i-1, j))
    """

  print(G)
  print(G.edges)
  l = nx.shortest_path_length(G, source=start, target=end)
  path = nx.shortest_path(G, start, end)
  print(path, len(path))
  
  print(l)