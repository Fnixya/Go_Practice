import numpy as np
import matplotlib.pyplot as plt

c = 1
L = 7

# Abreviation for the sines used on the Fourier coefficients
def b_n_sine(x, n):
    return np.sin((x * np.pi * n) / L)

# Fourier coefficients for the given f(x)
def b_n(n):
    # Compute b_n based on the results of part (a)
    mult = 294 / (4 * 7 * (np.pi ** 2) * (n ** 2))
    return 2 * mult * (3 * b_n_sine(3, n) - 2 * b_n_sine(5 / 2, n) - 2 * b_n_sine(9 / 2, n) + b_n_sine(5, n))

# Fourier series solution to the wave equation
def u(x, t, N=50):
    solution = np.zeros_like(x)
    for n in range(1, N + 1):
        solution += b_n(n) * np.cos(n * np.pi * t / L) * np.sin(n * np.pi * x / L)
    return solution

if __name__ == '__main__':
    # Plot the solution at different times
    x = np.linspace(0, 7, 500)
    t_values = [3, 4, 5]

    plt.figure(figsize=(10, 6))
    for t in t_values: 
        plt.plot(x, u(x, t, 500), label=f't = {t}')

    plt.title('Solution of the Wave Equation')
    plt.xlabel('x')
    plt.ylabel('u(x, t)')
    plt.legend()
    plt.grid(True)
    plt.show()