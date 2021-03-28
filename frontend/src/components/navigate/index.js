import { useEffect, useState } from "react";
import { useParams } from "react-router";
import axiosInstance from "../../axiosInstance";
import NoLinkFound from "../notfound";

function Navigate() {
  let { url } = useParams();
  const [link, setLink] = useState(false);
  const [serverError, setServerError] = useState(false);

  useEffect(() => {
    axiosInstance
      .get(`/short/${url}`)
      .then((res) => {
        const { status, originalURL } = res.data;
        if (status) {
          window.location.href = originalURL;
          return;
        }
        setLink(true);
      })
      .catch((err) => {
        setLink(true);
        setServerError(true);
      });
  }, [url]);

  return link && <NoLinkFound serverError={serverError} />;
}

export default Navigate;
