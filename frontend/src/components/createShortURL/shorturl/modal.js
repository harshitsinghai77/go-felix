import { Modal } from "antd";

function error(title, content) {
  Modal.error({ title, content });
}

export default error;
