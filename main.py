import nltk
import codecs
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords
#with open('three.txt', 'r') as f2:
with codecs.open('three.txt', 'r', "utf-8-sig") as f2:
    data = f2.read()
    print(data)
tokens = word_tokenize(data)
text = nltk.Text(tokens)

sr= stopwords.words('english')
clean_tokens = tokens[:]
for token in tokens:
    if token in stopwords.words('english'):
        
        clean_tokens.remove(token)
freq = nltk.FreqDist(clean_tokens)
for key,val in freq.items():
    key = key.encode("UTF-8", "replace")
    print(str(key) + ':' + str(val))
freq.plot(20, cumulative=False)
