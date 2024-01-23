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
        <label className="text-cyan-50 font-light font-mono pt-2 mx-4">
          Shorten a long URL
        </label>
        <input
          className="rounded-lg h-16 w-64 sm:w-96 text-left pl-4 mt-4 mb-6 mx-4 bg-slate-900 border-2 text-slate-50"
          onChange={handleChange}
          value={shortUrl || url}
          type="url"
          id="url"
          name="url"
          autoComplete="off"
          required
          placeholder="https://www.example.com"
        />

        {!shortUrl && (
          <Button>
            <p className="font-semibold font-mono">Send it yo</p>
          </Button>
        )}
      </form>
      {shortUrl && (
        <div className="flex justify-between items-center">
          <Button onClick={copyURL}>
            <p className="font-semibold font-mono text-center">Copy it yo</p>
          </Button>
          <div className="mt-5 h-14 bg-gradient-to-t from-transparent from-0% via-violet-300 via-50% to-transparent to-100% w-0.5"></div>
          <Button onClick={resetForm}>
            <p className="font-semibold font-mono text-center">Another Go</p>
          </Button>
        </div>
      )}
    </>
  );
}
