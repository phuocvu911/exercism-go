export function age(planet: string, seconds: number): number {
  const multiplier = orbitalPeriods.get(planet)
  if (!multiplier) throw new Error("invalid input")
  const earthAge = seconds/ (365.25*24*3600)
  return Math.round(earthAge*100/multiplier)/100
}

const orbitalPeriods: Map<string, number> = new Map([
  ["mercury", 0.2408467],
  ["venus", 0.61519726],
  ["earth", 1.0],
  ["mars", 1.8808158],
  ["jupiter", 11.862615],
  ["saturn", 29.447498],
  ["uranus", 84.016846],
  ["neptune", 164.79132],
]);

