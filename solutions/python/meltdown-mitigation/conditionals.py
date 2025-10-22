def is_criticality_balanced(temperature, neutrons_emitted):
    return temperature < 800 and neutrons_emitted > 500 and temperature * neutrons_emitted < 500000


def reactor_efficiency(voltage, current, theoretical_max_power):
    levels = ((80, "green"), (60, "orange"), (30, "red"), (0, "black"))
    efficiency = 100 * (voltage * current) / theoretical_max_power
    for level, color in levels:
        if efficiency >= level:
            return color
            

def fail_safe(temperature, neutrons_produced_per_second, threshold):
    mix = temperature * neutrons_produced_per_second
    print(neutrons_produced_per_second)
    if mix < threshold * .9:
        return 'LOW'
    if mix > threshold * 1.1:
        return 'DANGER'
    return 'NORMAL'