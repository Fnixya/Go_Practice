import numpy as np
import matplotlib.pyplot as plt
from matplotlib.widgets import Button, Slider


# The parametrized function to be plotted
def f(t, amplitude, frequency):
    return amplitude * np.sin(2 * np.pi * frequency * t)

# t = np.linspace(0, 1, 1000)

# Define initial parameters
# init_amplitude = 5
# init_frequency = 3




# The function to be called anytime a slider's value changes
def update(val):
    line.set_ydata(u(x, val, 500))
    fig.canvas.draw_idle()

def reset(event):
    freq_slider.reset()
    # amp_slider.reset()


c = 1
L = 7


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
    # t_values = [3, 4, 5]
    # for t in t_values: 
    #     ax.plot(x, u(x, t, 500), label=f't = {t}')

    # Create the figure and the line that we will manipulate
    fig, ax = plt.subplots(figsize=(10, 6))
    # plt.figure(figsize=(10, 6))
    line, = ax.plot(x, u(x, 3, 500), lw=2)
    # ax.set_xlabel('Time [s]')

    # adjust the main plot to make room for the sliders
    fig.subplots_adjust(left=0.25, bottom=0.25)

    # Make a horizontal slider to control the frequency.
    axfreq = fig.add_axes([0.25, 0.1, 0.65, 0.03])
    freq_slider = Slider(
        ax=axfreq,
        label='Time [t]',
        valmin=0,
        valmax=14,
        valinit=3,
    )

    # Make a vertically oriented slider to control the amplitude
    # axamp = fig.add_axes([0.1, 0.25, 0.0225, 0.63])
    # amp_slider = Slider(
    #     ax=axamp,
    #     label="Amplitude",
    #     valmin=0,
    #     valmax=10,
    #     valinit=init_amplitude,
    #     orientation="vertical"
    # )

    # Buttons _________________________________________________________________
    # register the update function with each slider
    freq_slider.on_changed(update)
    # amp_slider.on_changed(update)

    # Create a `matplotlib.widgets.Button` to reset the sliders to initial values.
    resetax = fig.add_axes([0.8, 0.025, 0.1, 0.04])
    button = Button(resetax, 'Reset', hovercolor='0.975')
    button.on_clicked(reset)


    plt.title('Solution of the Wave Equation')
    plt.xlabel('x')
    plt.ylabel('u(x, t)')
    plt.legend()
    plt.grid(True)
    plt.show()