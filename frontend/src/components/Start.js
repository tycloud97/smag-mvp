import React, { useState } from "react";
import { CameraFeed } from "./camera-feed";
import H1 from "./H1";
import H2 from "./H2";
import { FaceSearchRequest } from "../protofiles/usersearch_pb";
import IGPost from "./IGPost";

function findFacesInImage(apiClient, onFindFaces) {
  return async file => {
    const reader = new FileReader();
    reader.onloadend = async () => {
      const dataUrl = reader.result;
      const base64Data = dataUrl.split(",")[1];

      const request = new FaceSearchRequest();
      request.setBase64encodedpicture(base64Data);
      const response = await apiClient.searchSimilarFaces(request);
      const faces = response.getFacesList().map(protoFace => ({
        postId: protoFace.getPostId(),
        x: protoFace.getX(),
        y: protoFace.getY(),
        width: protoFace.getWidth(),
        height: protoFace.getHeight(),
        fullImageSrc: protoFace.getFullImageSrc()
      }));

      onFindFaces(faces);
    };
    reader.readAsDataURL(file);
  };
}

function Start({ apiClient, faceHits, addFaceHits }) {
  const onFileSubmit = findFacesInImage(apiClient, addFaceHits);

  return (
    <div className="body">
      <div className="column-full">
        <H1>
          Are you aware that wherever you are recorded, your identity can be
          found?
        </H1>
        <CameraFeed onFileSubmit={onFileSubmit} />
        <H2>Give us your smile face :)</H2>
      </div>
    </div>
  );
}

export default Start;
