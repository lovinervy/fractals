import matplotlib.pyplot as plt


g2 = 0
g3 = 4

# Качество изображения
img_x = 1000
img_y = 1000

# область отображение
xi = -2
xf = 2
yi = -2
yf = 2

iterations = 8 # N оттенков серого  

results = {} 

for y in range(img_y):
    zy = y * (yf - yi) / (img_y - 1)  + yi
    for x in range(img_x):
        zx = x * (xf - xi) / (img_x - 1)  + xi
        z = zx + zy * 1j
        for i in range(iterations):
            if abs(z) > 2:
                break
            z = (z**4 + 0.5*g2*z*z + 2*g3*z + 1/16*g2**2)/ \
                (4*z**3 - g2*z - g3)
        if i not in results:
            results[i] = [[], []]
        results[i][0].append(x)
        results[i][1].append(y)

for i, (xli, yli) in results.items():
    gray = 1.0 - i / iterations
    plt.plot(xli, yli, '.', color=(gray, gray, gray))

plt.show()  