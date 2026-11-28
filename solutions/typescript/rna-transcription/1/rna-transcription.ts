export function toRna(dna: string): string {
  return [...dna].map((x) => {
    const res = COMPLEMENTS[x]
    if (!res){
      throw new Error("Invalid input DNA.")
    }
    return res
  }). join("")
}

const COMPLEMENTS: Record<string, string> = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};
