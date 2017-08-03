
allwords = () ->
  line = io.read!
  pos = 1
  () ->
    while line do
      w,e = string.match(line, "(%w+[,;.:]?)()", pos)
      if w then
        pos = e
        return w
      else
        line = io.read!
        pos = i

    nil

prefix = (w1, w2) -> w1.." "..w2

statetab = {}

insert = (prefix, value) ->
  list = statetab[prefix]
  if not list then statetab[prefix] = value else list[#list + 1] = value

MAXGEN = 200
NOWORD = "\n"

-- build table
w1, w2 = NOWORD, NOWORD

for nextword in allwords! do
  insert(prefix(w1,w2), nextword)
  w1 = w2
  w2 = nextword
insert(prefix(w1,w2), NOWORD)

-- generate the text

w1 = NOWORD
w2 = NOWORD
for i = 1, MAXGEN do
  list = statetab[prefix(w1,w2)]
  -- choose a random item from list
  r =  math.random(#list)
  nextword = list[r]
  if nextword == NOWORD then return
  io.write nextword, " "
  w1 = w2
  w2 = nextword



