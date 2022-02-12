const Accounts = (): JSX.Element => {
  return (
    <table className="m-12 table-auto">
      <thead>
        <tr>
          <th>Username</th>
          <th>Email</th>
          <th>Created</th>
          <th>Updated</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>The Sliding Mr. Bones (Next Stop, Pottersville)</td>
          <td>Malcolm Lockyer</td>
          <td>1961</td>
          <td>1961</td>
        </tr>
        <tr>
          <td>Witchy Woman</td>
          <td>The Eagles</td>
          <td>1972</td>
          <td>1972</td>
        </tr>
        <tr>
          <td>Shining Star</td>
          <td>Earth, Wind, and Fire</td>
          <td>1975</td>
          <td>1975</td>
        </tr>
      </tbody>
    </table>
  );
};

export default Accounts;
