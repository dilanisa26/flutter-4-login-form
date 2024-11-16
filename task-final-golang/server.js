// server.js
const express = require('express');
const cors = require('cors');
const app = express();
const port = 8080;

// Allow Cross-Origin requests
app.use(cors({
    origin: 'http://localhost:5173',  // URL frontend
    credentials: true  // Allow sending cookies or credentials
}));

// Dummy account data
const accounts = [
  { id: 1, username: 'user1', email: 'user1@example.com', balance: 1000 },
  { id: 2, username: 'user2', email: 'user2@example.com', balance: 2000 },
  { id: 3, username: 'user3', email: 'user3@example.com', balance: 3000 },
  { id: 4, username: 'user4', email: 'user4@example.com', balance: 4000 }
];

// Define the /account/list endpoint
app.get('/account/list', (req, res) => {
  res.json(accounts);
});

// Define the /account/balance endpoint
app.get('/account/balance', (req, res) => {
  const totalBalance = accounts.reduce((sum, account) => sum + account.balance, 0);
  const averageBalance = accounts.length > 0 ? totalBalance / accounts.length : 0;
  res.json({ totalBalance, averageBalance });
});

// Start the server
app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});



// Contoh fungsi untuk mendapatkan total balance dan rata-rata balance
async function calculateTotalBalance() {
  // Gantilah ini dengan logika untuk mengambil total balance dari database Anda
  return 10000; // Contoh nilai total balance
}

async function getTotalAccounts() {
  // Gantilah ini dengan logika untuk mengambil jumlah total akun dari database Anda
  return 4; // Contoh jumlah akun
}

// Endpoint untuk mendapatkan balance
app.get('/account/balance', async (req, res) => {
  try {
    const totalBalance = await calculateTotalBalance();
    const totalAccounts = await getTotalAccounts();
    const averageBalance = totalAccounts > 0 ? totalBalance / totalAccounts : 0;

    res.json({ totalBalance, averageBalance });
  } catch (error) {
    console.error('Error fetching balance:', error);
    res.status(500).json({ message: 'Internal Server Error' });
  }
});

app.listen(8080, () => {
  console.log('Server is running on http://localhost:8080');
});
