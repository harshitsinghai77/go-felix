import { useState } from "react";
import { Link } from "react-router-dom";
import { Card, Input, Divider, Button, Spin } from "antd";
import ModalError from "./modal";
import styles from "./index.module.css";
import axiosInstance from "../../../axiosInstance";

import "antd/lib/spin/style/css";
import "antd/lib/card/style/css";
import "antd/lib/input/style/css";
import "antd/lib/select/style/css";
import "antd/lib/divider/style/css";
import "antd/lib/modal/style/css";

function App() {
  const [loader, setLoader] = useState(false);
  const [originalURL, setOriginalURL] = useState("");
  const [shortURL, setShortURL] = useState("");
  const [alreadyExists, setAlreadyExists] = useState(false);

  const onSubmit = () => {
    setLoader(true);
    if (originalURL) {
      const postBody = {
        original_url: originalURL,
      };

      axiosInstance
        .post("/short", postBody)
        .then((res) => {
          const { shortUrl, alreadyExists } = res.data;
          setAlreadyExists(alreadyExists);
          setShortURL(shortUrl);
          setLoader(false);
        })
        .catch((err) => {
          ModalError(
            "It looks like our server has trouble processing your request. Hold on while we fix the issue."
          );
          setLoader(false);
        });
    } else {
      ModalError("URL Cannot be empty", "Please enter a valid url.");
      setLoader(false);
    }
  };

  return (
    <div className={styles.app_magin}>
      <Card className={styles.cardstyle} bodyStyle={{ padding: "2em" }}>
        <h1 className={styles.title}>Felix</h1>
        <p className={styles.text}>Enter Long URL</p>
        <Input
          allowClear
          addonBefore="Enter URL"
          // defaultValue="mysite"
          value={originalURL}
          className={styles.inputtext}
          onChange={(e) => {
            if (e.target.value === "") {
              setAlreadyExists(false);
              setShortURL("");
            }
            setOriginalURL(e.target.value);
          }}
        />
        <p className={styles.marginBottom}></p>

        <Button onClick={onSubmit} disabled={!originalURL}>
          Generate Short Link
        </Button>
        {loader && (
          <div className={styles.spinner}>
            <Spin size="large" />
          </div>
        )}
        {shortURL && (
          <>
            <Divider />
            <p className={styles.generatedLinkText}>
              {alreadyExists ? "Link Already Exists" : "Generated Link "}
              &nbsp;
              <Link to={`/${shortURL}`}>
                <span className={styles.hyperlink}>
                  {process.env.REACT_APP_SHORTNER_LINK}/{shortURL}
                </span>
              </Link>
            </p>
          </>
        )}
      </Card>
    </div>
  );
}

export default App;
