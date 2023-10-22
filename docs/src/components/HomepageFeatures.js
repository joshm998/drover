import React from "react";
import clsx from "clsx";
import styles from "./HomepageFeatures.module.css";

const FeatureList = [
  {
    title: "Lightweight",
    Svg: require("../../static/img/lightweight.svg").default,
    description: (
      <>
        Designed from the ground up to run on almost every piece of hardware including the Pi Zero.
      </>
    ),
  },
  {
    title: "Unified at Last",
    Svg: require("../../static/img/unified-handshake.svg").default,
    description: (
      <>
        Use all of your 3D printers in one place, supporting Prusa, Tiertime, and more...
      </>
    ),
  },
  {
    title: "API-first",
    Svg: require("../../static/img/api-puzzle.svg").default,
    description: (
      <>
        Don't want to use our user inteface? No problem, our APIs can be used standalone.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center">
        <Svg className={styles.featureSvg} alt={title} />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
