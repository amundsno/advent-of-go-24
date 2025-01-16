from math import gcd, floor, ceil
import re

def parse_input(path):
    rows = re.findall('(\d+).*?(\d+).*?(\d+).*?(\d+).*?(\d+).*?(\d+)',
                 open(path).read(), re.S)
    
    return rows


PRICE_A = 3
PRICE_B = 1

def lde(a, b, c):
    # From https://new.math.uiuc.edu/public348/python/diophantus.html
    q, r = divmod(a, b)
    if r == 0:
        return ([0, c/b])
    else:
        sol = lde(b, r, c)
        u = sol[0]
        v = sol[1]
        return ([v, u-q*v])


def cost_to_price(row):
    ax, ay, bx, by, tx, ty = map(int, row)
    
    det = ax*by - bx*ay
    if det != 0:
        # Case 1: Only one possible solution
        aDet = tx*by - ty*bx
        bDet = ty*ax - tx*ay
        
        if aDet % det == 0 and bDet % det == 0:
            # The solution is valid only A and B are integers
            A, B = aDet//det, bDet//det
            return PRICE_A*A + PRICE_B*B
        
        return -1
    
    detAug = ax*ty - tx*ay
    if detAug == 0 and tx % gcd(ax, bx) != 0:
        # Case 2: Many possible solutions, but none are valid
        return -1
    
    # Case 3: Many possible solutions, but only one is optimal
    # Find one solution to the LDE: A(ax) + B(bx) = tx
    A0, B0 = lde(ax, bx, tx)
    
    # General solutions are of the form: A = A0 + k*bx, B = B0 - k*ax
    # Select the k that minimizes the cost inefficient button
    k = [ceil(-A0/bx), floor(B0/ax)]
    k = max(k[0], k[1]) if ax/bx > PRICE_A else min(k[0], k[1])
    
    A = A0+k*bx
    B = B0-k*ax
    
    if A < 0 or B < 0:
        # Invalid solution, despite selecting optimal k
        return -1
    
    return PRICE_A*A + PRICE_B*B


if __name__ == "__main__":
    sum = 0
    rows = parse_input('input.txt')
    for row in rows:
        cost = cost_to_price(row)
        if cost > 0:
            sum += cost

    print(sum)
    
    # Hypothetical cases
    print(cost_to_price((4, 4, 3, 3, 5, 5))) # Should not be possible
    print(cost_to_price((1,1, 2,2, 5,5))) # Should be optimal with 1*A + 2*B = 5 tokens
    print(cost_to_price((7, 7, 2, 2, 20, 20))) # Should be optimal with 2*A + 3*B = 9 tokens
    