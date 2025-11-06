import React, { useState, useEffect } from "react";
import { motion, AnimatePresence } from "framer-motion";
import ExperienceCard from "./ExperienceCard";

const variants = {
  enter: (direction) => ({
    x: direction > 0 ? "100%" : "-100%",
    opacity: 0,
  }),
  center: {
    x: 0,
    opacity: 1,
  },
  exit: (direction) => ({
    x: direction < 0 ? "100%" : "-100%",
    opacity: 0,
  }),
};

export default function ExperienceCarousel({ jobs, jobIndex, direction, onPaginate, onDragEnd }) {
  
    return (
        <div className="experience-carousel-container">
          <div className="experience-sizer" aria-hidden="true">
            <ExperienceCard job={jobs[jobIndex]} />
          </div>
          <div className="experience-animator">
            <AnimatePresence initial={false} custom={direction}>
              <motion.div
                key={jobIndex}
                className="experience-slide"
                custom={direction}
                variants={variants}
                initial="enter"
                animate="center"
                exit="exit"
                transition={{
                  x: { type: "spring", stiffness: 300, damping: 30 },
                  opacity: { duration: 0.2 },
                }}
                drag="x"
                dragConstraints={{ left: 0, right: 0 }}
                dragElastic={0.1}
                onDragEnd={onDragEnd}
              >
                <ExperienceCard job={jobs[jobIndex]} />
              </motion.div>
            </AnimatePresence>
          </div>
        </div>
      );
}