import { useEffect, useState } from "react";

function App() {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("http://localhost:8080/")
      .then((res) => res.json())
      .then((data) => {
        setUsers(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Error fetching data:", err);
        setLoading(false);
      });
  }, []);

  return (
    <div
      style={{
        minHeight: "100vh",
       // backgroundColor: "#eef2f7",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        fontFamily: "Segoe UI, Arial, sans-serif",
        padding: "20px",
      }}
    >
      <div
        style={{
          marginLeft:"700px",
          //width: "100%",
          //maxWidth: "700px",
          backgroundColor: "#ffffff",
          borderRadius: "16px",
          boxShadow: "0 15px 40px rgba(0,0,0,0.08)",
          padding: "30px",
        }}
      >
        <h1
          style={{
            textAlign: "center",
            color: "#2c3e50",
            marginBottom: "25px",
          }}
        >
          📊 Excel Data Viewer
        </h1>

        {loading ? (
          <p
            style={{
              textAlign: "center",
              fontSize: "18px",
              color: "#666",
            }}
          >
            Loading data...
          </p>
        ) : users.length === 0 ? (
          <p
            style={{
              textAlign: "center",
              fontSize: "18px",
              color: "#666",
            }}
          >
            No data found
          </p>
        ) : (
          <div
            style={{ display: "flex", flexDirection: "column", gap: "12px" }}
          >
            {users.map((d) => (
              <div
                key={d.id}
                style={{
                  backgroundColor: "#f8fafc",
                  padding: "16px 20px",
                  borderRadius: "12px",
                  border: "1px solid #e2e8f0",
                  display: "flex",
                  flexDirection: "column",
                  gap: "6px",
                  textAlign: "center",
                }}
              >
                <div>
                  <strong style={{ color: "#1f2937" }}>ID:</strong>{" "}
                  <span style={{ color: "#4b5563" }}>{d.id}</span>
                </div>
                <div>
                  <strong style={{ color: "#1f2937" }}>Name:</strong>{" "}
                  <span style={{ color: "#4b5563" }}>{d.name}</span>
                </div>
                <div>
                  <strong style={{ color: "#1f2937" }}>Email:</strong>{" "}
                  <span style={{ color: "#4b5563" }}>{d.email}</span>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
