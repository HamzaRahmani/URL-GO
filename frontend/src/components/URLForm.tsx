import { type FormEvent, useState } from "react";
import Button from "./Button.tsx";

export default function Form() {
  const [responseMessage, setResponseMessage] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [url, setUrl] = useState("");

  function submit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    try {
      fetch(`${import.meta.env.PUBLIC_URLAPI}/url`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          url: url,
        }),
      }).then(async (response) => {
        if (response.ok) {
          const data = await response.json();
          console.log(data);
          if (data.hash) {
            setShortUrl(`${import.meta.env.PUBLIC_URLAPI}/${data.hash}`);
          }
          return;
        }

        throw new Error("Something went wrong, try again later");
      });
    } catch (error) {
      console.log(error);
    }
  }

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(event.target.value);
  };

  const resetForm = () => {
    setShortUrl("");
    setUrl("");
  };

  const copyURL = () => {
    navigator.clipboard.writeText(shortUrl);
  };

  return (
    // TODO: Break down into smaller components
    <>
      <form onSubmit={submit} className="flex flex-col">
        <label className="text-cyan-50 font-light font-mono pt-2">
          Shorten a long URL
        </label>
        <input
          className="rounded-lg h-10 w-96 text-left pl-1 mt-4 mb-6"
          onChange={handleChange}
          value={shortUrl || url}
          type="url"
          id="url"
          name="url"
          required
        />
        {!shortUrl && (
          <Button>
            <p className="font-semibold font-mono px-2">Send it yo</p>
          </Button>
        )}
      </form>
      {shortUrl && (
        <div className="flex justify-between">
          <Button onClick={copyURL}>
            <p className="font-semibold font-mono text-center px-2">
              Copy it yo
            </p>
          </Button>
          <Button onClick={resetForm}>
            <p className="font-semibold font-mono text-center px-2">
              Another Go
            </p>
          </Button>
        </div>
      )}
    </>
  );
}
