import { type FormEvent, useState } from "react";
import Button from "./Button.tsx";

export default function Form() {
  const [responseMessage, setResponseMessage] = useState("");
  const [url, setUrl] = useState("");

  function submit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    // TODO: load url from env
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
          if (data.hash) {
            setResponseMessage(data.hash);
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
    // TODO: Move to a new component
    // TODO: When successful render a component with short URL + copy option
    <form onSubmit={submit} className="flex flex-col">
      <label className="text-cyan-50">Shorten a long URL</label>
      <input
        className="rounded-lg h-10 w-96 text-left pl-1 my-4"
        onChange={handleChange}
        value={url}
        type="url"
        id="url"
        name="url"
        autoComplete="url"
        required
      />
      <Button>
        <p className="font-semibold font-mono">Send it yo</p>
      </Button>
      {responseMessage && <p>{responseMessage}</p>}
    </form>
  );
}
