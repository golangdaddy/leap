"use client"

export default function Loading(props) {

    var size = 120
    switch (props.small) {
        case true:
            size = 60
    }

    return <>
      <style jsx>{`
          .loader {
            border: 16px solid #f3f3f3; 
            border-top: 16px solid #3094da;
            border-radius: 50%;
            animation: spin 1s linear infinite;
            margin: auto;
          }
  
          @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
          }
      `}</style>
      <div className="loader" style={{width:size+"px", height:size+"px"}}></div> 
    </>
  }