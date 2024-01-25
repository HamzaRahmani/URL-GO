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
    <div className="flex flex-col justify-start items-center h-full">
      <form
        onSubmit={submit}
        className="flex flex-col justify-evenly items-center h-full w-5/6"
      >
        <label className="text-cyan-50 font-light font-mono pt-2 ml-0 w-5/6">
          Shorten a long URL
        </label>
        <input
          className="rounded-lg bg-slate-900 border-2 text-slate-50 h-16 pl-4 w-5/6"
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
          <Button className="w-5/6">
            <p className="font-semibold font-mono">Send it yo</p>
          </Button>
        )}
        {shortUrl && (
          <div className="flex justify-between items-center w-5/6">
            <Button onClick={copyURL}>
              <p className="font-semibold font-mono text-center">Copy it yo</p>
            </Button>
            <div className="bg-gradient-to-t from-transparent from-0% via-violet-300 via-50% to-transparent to-100% w-0.5 mt-5 h-14"></div>
            <Button onClick={resetForm}>
              <p className="font-semibold font-mono text-center">Another Go</p>
            </Button>
          </div>
        )}
      </form>
    </div>
  );
}
