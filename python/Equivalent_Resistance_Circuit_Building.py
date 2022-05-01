import sys
import math

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

def get_idx_end_block(c):
    try:
        i = c.index(")")
    except:
        i = -1
    try:
        j = c.index("]")
    except:
        j = -1

    if (i <= j and i != -1) or j == -1:
        return i
    else:
        return j

def get_idx_start_block(circuit_list, idx_end):
    i = idx_end
    c = "(" if circuit_list[idx_end] == ")" else "["
    while i >= 0:
        if circuit_list[i] != c:
            i -= 1
        else:
            break
    return i

def get_r_series(resistors, l):
    return sum([resistors[i] for i in l])

def get_r_parallel(resistors, l):
    s = sum([1.0/resistors[i] for i in l])
    return 1.0/s

resistors = dict()

n = int(input())
for i in range(n):
    inputs = input().split()
    name = inputs[0]
    r = int(inputs[1])
    resistors[name] = r

circuit = input()
circuit_list = circuit.split()

idx_end = get_idx_end_block(circuit_list)
while idx_end != -1:
    idx_start = get_idx_start_block(circuit_list, idx_end)
    if circuit_list[idx_start] == "(":
        r = get_r_series(resistors, circuit_list[idx_start+1:idx_end])
    else:
        r = get_r_parallel(resistors, circuit_list[idx_start+1:idx_end])

    resistors[str(r)] = r
    circuit_list[idx_start] = str(r)
    for i in range(idx_end-idx_start):
        circuit_list.pop(idx_start+1)
    idx_end = get_idx_end_block(circuit_list)

print("{:.1f}".format(float(circuit_list[0])))
