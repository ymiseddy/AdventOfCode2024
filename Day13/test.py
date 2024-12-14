import sympy as sp


if __name__ == '__main__':
    a, b, x1, x2, y1, y2, xf, yf = sp.symbols('a b x1 x2 y1 y2 xf yf')

    eqa = sp.Eq(x1*a + x2*b, xf)
    eqb = sp.Eq(y1*a + y2*b, yf)
    sol = sp.solve([eqa, eqb], [a, b])
    print(sol)
