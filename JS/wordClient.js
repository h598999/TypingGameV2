export async function fetchWords(){
  try {
    const response = await fetch("/words")
    if (!response.ok){
      throw new Error("WordClient failed, couldnt fetch the words");
    }
    const data = await response.json();
    return data;
  } catch(error){
    console.error("Error fetching the words: " ,error);
    return [];
  }
} 

export async function fetchLorem(){
  try {
    const response = await fetch("/lorem_ipsum")
    if (!response.ok){
      throw new Error("WordClient failed, couldnt fetch the words");
    }
    const data = await response.json();
    return data;
  } catch(error){
    console.error("Error fetching the words: " ,error);
    return null; 
  }
} 
