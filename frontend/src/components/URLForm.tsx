import { type FormEvent, useState } from "react";

export default function Form() {
  const [responseMessage, setResponseMessage] = useState("");
  const [url, setUrl] = useState("");

  function submit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    fetch("http://localhost:5050/url", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        url: url,
      }),
    })
      .then(async (response) => {
        if (response.ok) {
          const data = await response.json();
          console.log(data);
          if (data.message) {
            setResponseMessage(data.message);
          }
          return;
        }

        throw new Error("Something went wrong, try again later");
      })
      .catch((error) => {
        console.log(error);
      });
  }

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(event.target.value);
  };

  return (
    <form onSubmit={submit}>
      <label>Enter a URL</label>
      <input
        onChange={handleChange}
        value={url}
        type="url"
        id="url"
        name="url"
        autoComplete="url"
        required
      />
      <button>Send</button>
      {responseMessage && <p>{responseMessage}</p>}
    </form>
  );
}
