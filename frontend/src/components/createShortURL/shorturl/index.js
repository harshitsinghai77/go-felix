import { useState } from "react";
import { Link } from "react-router-dom";
import { Card, Input, Select, Divider } from "antd";
import ModalError from "./modal";
import styles from "./index.module.css";
import axiosInstance from "../../../axiosInstance";

// /button/style/index.less
import "antd/lib/card/style/css";
import "antd/lib/input/style/css";
import "antd/lib/select/style/css";
import "antd/lib/divider/style/css";
import "antd/lib/modal/style/css";

const { Option } = Select;

function App() {
  const [originalURL, setOriginalURL] = useState("");
  const [expiryDateType, setExpiryDateType] = useState(null);
  const [expiresAfter, setExpiresAfter] = useState(null);
  const [shortURL, setShortURL] = useState("");
  const [alreadyExists, setAlreadyExists] = useState(false);

  const onExpireChange = (e) => {
    const onlyNumb = e.target.value.replace(/\D/, "");
    setExpiresAfter(onlyNumb);
  };

  const onSubmit = () => {
    if (originalURL) {
      const postBody = {
        original_url: originalURL,
        expiry: expiryDateType !== "none",
        expires_after: expiresAfter,
        expiry_date_type: expiryDateType,
      };

      if (postBody.expiry && !postBody.expires_after) {
        ModalError(
          "'Expire after' cannot be blank",
          "Either set expire after to None or enter value"
        );
        return;
      }
      axiosInstance
        .post("/short", postBody)
        .then((res) => {
          const { shortUrl, alreadyExists } = res.data;
          setAlreadyExists(alreadyExists);
          setShortURL(shortUrl);
        })
        .catch((err) => {
          ModalError(
            "It looks like our server has trouble processing your request. Hold on while we fix the issue."
          );
        });
    } else {
      ModalError("URL Cannot be empty", "Please enter a valid url.");
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
        <p className={styles.text}>Enter Long URL</p>
        <Input
          addonBefore="Expire after"
          addonAfter={
            <Select
              defaultValue="none"
              onChange={(el) => setExpiryDateType(el)}
              className="select-after"
            >
              <Option value="none">None</Option>
              <Option value="min">minute</Option>
              <Option value="hour">hour</Option>
              <Option value="day">day</Option>
            </Select>
          }
          value={expiresAfter}
          onChange={onExpireChange}
          disabled={expiryDateType === "none"}
          className={styles.inputtext}
        />

        <p className={styles.marginBottom}></p>

        <button className={styles.button5} onClick={onSubmit}>
          Generate Short Link
        </button>

        {shortURL && (
          <>
            <Divider />
            <p className={styles.generatedLinkText}>
              {alreadyExists ? "Link Already Exists" : "Generated Link "}
              &nbsp;
              <Link to={`/${shortURL}`}>
                <span className={styles.hyperlink}>
                  localhost:3000/{shortURL}
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
