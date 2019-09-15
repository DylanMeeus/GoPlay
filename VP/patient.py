import sys

class doc:
    def __init__(self, name):
        self.name = name
        self.properties = []

    def __str__(self):
        return str(self.name) + ": " +  ",".join(map(lambda k: str(k),self.properties))

    def __repr__(self):
        return str(self.name)


def solve(docs, vw):
    # match the best doctor to a patient 
    docmatches = {}
    docqs = {}
    for d in docs:
        M = (list(filter(lambda k: k in vw, d.properties))) 
        docmatches[d] = M
        q = list(filter(lambda k: k not in vw, d.properties))
        docqs[d] = q
    print(docmatches)
    print(docqs)
    while len(docs) > 1:
        # try to remove a doctor
        to_remove = []
        for doc1, es in docmatches.items():
            for doc2, es2 in docmatches.items():
                if doc1 == doc2:
                    if len(es2) == 0 and len(es) != 0:
                        to_remove.append(doc2)
                if len(es) > len(es2):
                    to_remove.append(doc2)
        docs = list(filter(lambda k: k not in to_remove, docs))
    return docs[0]




if __name__ == '__main__':
    lines = []
    for line in sys.stdin:
        lines.append(line)

    ptr = 1
    doctors = None
    
    while ptr < len(lines):
        if doctors == None:
            docs = []
            doctors = int(lines[ptr])
            ptr += 1
            for i in range(doctors):
                docline = lines[ptr]
                parts = list(map(lambda k: int(k), docline.replace("\n","").split(" ")))
                d = doc(parts[0])
                if len(parts) > 1:
                    d.properties = parts[1:]
                ptr += 1
                docs.append(d)
                print(d)
            # now the line is the patient line
            patient = lines[ptr]
            voorkeuren = list(map(lambda k: int(k), patient.replace("\n","").split(" ")))[1:]
            solve(docs, voorkeuren)




        ptr += 1
